package testdata

var ServerNoPayloadNoResultHandlerConstructorCode = `// NewMethodNoPayloadNoResultHandler creates a HTTP handler which loads the
// HTTP request and calls the "ServiceNoPayloadNoResult" service
// "MethodNoPayloadNoResult" endpoint.
func NewMethodNoPayloadNoResultHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeMethodNoPayloadNoResultResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "MethodNoPayloadNoResult")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceNoPayloadNoResult")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
`

var ServerNoPayloadNoResultWithRedirectHandlerConstructorCode = `// NewMethodNoPayloadNoResultHandler creates a HTTP handler which loads the
// HTTP request and calls the "ServiceNoPayloadNoResult" service
// "MethodNoPayloadNoResult" endpoint.
func NewMethodNoPayloadNoResultHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "MethodNoPayloadNoResult")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceNoPayloadNoResult")
		http.Redirect(w, r, "/redirect/dest", http.StatusMovedPermanently)
	})
}
`

var ServerPayloadNoResultHandlerConstructorCode = `// NewMethodPayloadNoResultHandler creates a HTTP handler which loads the HTTP
// request and calls the "ServicePayloadNoResult" service
// "MethodPayloadNoResult" endpoint.
func NewMethodPayloadNoResultHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMethodPayloadNoResultRequest(mux, decoder)
		encodeResponse = EncodeMethodPayloadNoResultResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "MethodPayloadNoResult")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ServicePayloadNoResult")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
`

var ServerPayloadNoResultWithRedirectHandlerConstructorCode = `// NewMethodPayloadNoResultHandler creates a HTTP handler which loads the HTTP
// request and calls the "ServicePayloadNoResult" service
// "MethodPayloadNoResult" endpoint.
func NewMethodPayloadNoResultHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest = DecodeMethodPayloadNoResultRequest(mux, decoder)
		encodeError   = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "MethodPayloadNoResult")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ServicePayloadNoResult")
		_, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		http.Redirect(w, r, "/redirect/dest", http.StatusMovedPermanently)
	})
}
`

var ServerNoPayloadResultHandlerConstructorCode = `// NewMethodNoPayloadResultHandler creates a HTTP handler which loads the HTTP
// request and calls the "ServiceNoPayloadResult" service
// "MethodNoPayloadResult" endpoint.
func NewMethodNoPayloadResultHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeMethodNoPayloadResultResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "MethodNoPayloadResult")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceNoPayloadResult")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
`

var ServerPayloadResultHandlerConstructorCode = `// NewMethodPayloadResultHandler creates a HTTP handler which loads the HTTP
// request and calls the "ServicePayloadResult" service "MethodPayloadResult"
// endpoint.
func NewMethodPayloadResultHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMethodPayloadResultRequest(mux, decoder)
		encodeResponse = EncodeMethodPayloadResultResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "MethodPayloadResult")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ServicePayloadResult")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
`

var ServerPayloadResultErrorHandlerConstructorCode = `// NewMethodPayloadResultErrorHandler creates a HTTP handler which loads the
// HTTP request and calls the "ServicePayloadResultError" service
// "MethodPayloadResultError" endpoint.
func NewMethodPayloadResultErrorHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMethodPayloadResultErrorRequest(mux, decoder)
		encodeResponse = EncodeMethodPayloadResultErrorResponse(encoder)
		encodeError    = EncodeMethodPayloadResultErrorError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "MethodPayloadResultError")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ServicePayloadResultError")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
`

var ServerSkipResponseBodyEncodeDecodeCode = `// NewMethodSkipResponseBodyEncodeDecodeHandler creates a HTTP handler which
// loads the HTTP request and calls the "ServiceSkipResponseBodyEncodeDecode"
// service "MethodSkipResponseBodyEncodeDecode" endpoint.
func NewMethodSkipResponseBodyEncodeDecodeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeMethodSkipResponseBodyEncodeDecodeResponse(encoder)
		encodeError    = EncodeMethodSkipResponseBodyEncodeDecodeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "MethodSkipResponseBodyEncodeDecode")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ServiceSkipResponseBodyEncodeDecode")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		o := res.(*serviceskipresponsebodyencodedecode.MethodSkipResponseBodyEncodeDecodeResponseData)
		defer o.Body.Close()
		if wt, ok := o.Body.(io.WriterTo); ok {
			if err := encodeResponse(ctx, w, res); err != nil {
				errhandler(ctx, w, err)
				return
			}
			n, err := wt.WriteTo(w)
			if err != nil {
				if n == 0 {
					if err := encodeError(ctx, w, err); err != nil {
						errhandler(ctx, w, err)
					}
				} else {
					if f, ok := w.(http.Flusher); ok {
						f.Flush()
					}
					panic(http.ErrAbortHandler) // too late to write an error
				}
			}
			return
		}
		// handle immediate read error like a returned error
		buf := bufio.NewReader(o.Body)
		if _, err := buf.Peek(1); err != nil && err != io.EOF {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
			return
		}
		if _, err := io.Copy(w, buf); err != nil {
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			panic(http.ErrAbortHandler) // too late to write an error
		}
	})
}
`
