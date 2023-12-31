// Code generated by goa v3.12.4, DO NOT EDIT.
//
// front HTTP server encoders and decoders
//
// Command:
// $ goa gen game-service/design

package server

import (
	"context"
	"errors"
	front "game-service/gen/front"
	frontviews "game-service/gen/front/views"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListCharactersResponse returns an encoder for responses returned by
// the front list-characters endpoint.
func EncodeListCharactersResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(frontviews.StoredCharacterCollection)
		enc := encoder(ctx, w)
		body := NewStoredCharacterResponseTinyCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeShowCharacterResponse returns an encoder for responses returned by the
// front show-character endpoint.
func EncodeShowCharacterResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*frontviews.StoredCharacter)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body any
		switch res.View {
		case "default", "":
			body = NewShowCharacterResponseBody(res.Projected)
		case "tiny":
			body = NewShowCharacterResponseBodyTiny(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowCharacterRequest returns a decoder for requests sent to the front
// show-character endpoint.
func DecodeShowCharacterRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id   string
			view *string
			err  error

			params = mux.Vars(r)
		)
		id = params["id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []any{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowCharacterPayload(id, view)

		return payload, nil
	}
}

// EncodeShowCharacterError returns an encoder for errors returned by the
// show-character front endpoint.
func EncodeShowCharacterError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *front.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewShowCharacterNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAddCharacterResponse returns an encoder for responses returned by the
// front add-character endpoint.
func EncodeAddCharacterResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddCharacterRequest returns a decoder for requests sent to the front
// add-character endpoint.
func DecodeAddCharacterRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body AddCharacterRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddCharacterRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddCharacterCharacter(&body)

		return payload, nil
	}
}

// EncodeAddCharacterError returns an encoder for errors returned by the
// add-character front endpoint.
func EncodeAddCharacterError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "name_taken":
			var res *front.NameTaken
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddCharacterNameTakenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateCharacterResponse returns an encoder for responses returned by
// the front update-character endpoint.
func EncodeUpdateCharacterResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeUpdateCharacterRequest returns a decoder for requests sent to the
// front update-character endpoint.
func DecodeUpdateCharacterRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateCharacterRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateCharacterRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewUpdateCharacterPayload(&body, id)

		return payload, nil
	}
}

// EncodeUpdateCharacterError returns an encoder for errors returned by the
// update-character front endpoint.
func EncodeUpdateCharacterError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "name_taken":
			var res *front.NameTaken
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateCharacterNameTakenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_found":
			var res *front.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateCharacterNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRemoveCharacterResponse returns an encoder for responses returned by
// the front remove-character endpoint.
func EncodeRemoveCharacterResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeRemoveCharacterRequest returns a decoder for requests sent to the
// front remove-character endpoint.
func DecodeRemoveCharacterRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewRemoveCharacterPayload(id)

		return payload, nil
	}
}

// EncodeRemoveCharacterError returns an encoder for errors returned by the
// remove-character front endpoint.
func EncodeRemoveCharacterError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *front.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewRemoveCharacterNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAddItemResponse returns an encoder for responses returned by the front
// add-item endpoint.
func EncodeAddItemResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddItemRequest returns a decoder for requests sent to the front
// add-item endpoint.
func DecodeAddItemRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body AddItemRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddItemRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddItemPayload(&body)

		return payload, nil
	}
}

// EncodeAddItemError returns an encoder for errors returned by the add-item
// front endpoint.
func EncodeAddItemError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "name_taken":
			var res *front.NameTaken
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddItemNameTakenResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_found":
			var res *front.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddItemNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRemoveItemResponse returns an encoder for responses returned by the
// front remove-item endpoint.
func EncodeRemoveItemResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeRemoveItemRequest returns a decoder for requests sent to the front
// remove-item endpoint.
func DecodeRemoveItemRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body RemoveItemRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateRemoveItemRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewRemoveItemPayload(&body)

		return payload, nil
	}
}

// EncodeRemoveItemError returns an encoder for errors returned by the
// remove-item front endpoint.
func EncodeRemoveItemError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *front.NotFound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewRemoveItemNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalFrontviewsStoredCharacterViewToStoredCharacterResponseTiny builds a
// value of type *StoredCharacterResponseTiny from a value of type
// *frontviews.StoredCharacterView.
func marshalFrontviewsStoredCharacterViewToStoredCharacterResponseTiny(v *frontviews.StoredCharacterView) *StoredCharacterResponseTiny {
	res := &StoredCharacterResponseTiny{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// unmarshalCharacterRequestBodyToFrontCharacter builds a value of type
// *front.Character from a value of type *CharacterRequestBody.
func unmarshalCharacterRequestBodyToFrontCharacter(v *CharacterRequestBody) *front.Character {
	res := &front.Character{
		Name:        *v.Name,
		Description: v.Description,
		Health:      *v.Health,
		Experience:  v.Experience,
	}

	return res
}

// unmarshalItemRequestBodyToFrontItem builds a value of type *front.Item from
// a value of type *ItemRequestBody.
func unmarshalItemRequestBodyToFrontItem(v *ItemRequestBody) *front.Item {
	res := &front.Item{
		Name:        *v.Name,
		Description: v.Description,
		Damage:      *v.Damage,
		Healing:     *v.Healing,
		Protection:  *v.Protection,
	}

	return res
}
