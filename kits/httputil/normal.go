package httputil

import (
	"bytes"
	"fmt"
	"github.com/XTeam-Wing/xkit/kits/strutils"
	"io"
	"net/http"
	"strings"

	"github.com/dsnet/compress/brotli"
	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zlib"
	"github.com/klauspost/compress/zstd"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// readNNormalizeRespBody performs normalization on the http response object.
// and fills body buffer with actual response body.
func readNNormalizeRespBody(rc *ResponseChain, body *bytes.Buffer) (err error) {
	response := rc.resp
	if response == nil {
		return fmt.Errorf("something went wrong response is nil")
	}
	// net/http doesn't automatically decompress the response body if an
	// encoding has been specified by the user in the request so in case we have to
	// manually do it.

	origBody := rc.resp.Body
	if origBody == nil {
		// skip normalization if body is nil
		return nil
	}
	// wrap with decode if applicable
	wrapped, err := wrapDecodeReader(response)
	if err != nil {
		wrapped = origBody
	}
	limitReader := io.LimitReader(wrapped, 2*MaxBodyRead)

	// read response body to buffer
	_, err = body.ReadFrom(limitReader)
	if err != nil {
		if strings.Contains(err.Error(), "gzip: invalid header") {
			// its invalid gzip but we will still use it from original body
			_, gErr := body.ReadFrom(origBody)
			if gErr != nil {
				return errors.Wrap(gErr, "could not read response body after gzip error")
			}
		} else if strutils.ContainsAnyI(err.Error(), "unexpected EOF", "read: connection reset by peer", "user canceled", "http: request body too large") {
			// keep partial body and continue (skip error) (add meta header in response for debugging)
			if response.Header == nil {
				response.Header = make(http.Header)
			}
			response.Header.Set("x-nuclei-ignore-error", err.Error())
			return nil
		} else {
			return errors.Wrap(err, "could not read response body")
		}
	}
	return nil
}

// wrapDecodeReader wraps a decompression reader around the response body if it's compressed
// using gzip or deflate.
func wrapDecodeReader(resp *http.Response) (rc io.ReadCloser, err error) {
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		rc, err = gzip.NewReader(resp.Body)
	case "deflate":
		rc, err = zlib.NewReader(resp.Body)
	case "br":
		rc, err = brotli.NewReader(resp.Body, nil)
	case "zstd":
		var zstdReader *zstd.Decoder
		zstdReader, err = zstd.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		rc = io.NopCloser(zstdReader)
	default:
		rc = resp.Body
	}
	if err != nil {
		return nil, err
	}
	// handle GBK encoding
	if isContentTypeGbk(resp.Header.Get("Content-Type")) {
		rc = io.NopCloser(transform.NewReader(rc, simplifiedchinese.GBK.NewDecoder()))
	}
	return rc, nil
}

// isContentTypeGbk checks if the content-type header is gbk
func isContentTypeGbk(contentType string) bool {
	contentType = strings.ToLower(contentType)
	return strutils.ContainsAny(contentType, "gbk", "gb2312", "gb18030")
}
