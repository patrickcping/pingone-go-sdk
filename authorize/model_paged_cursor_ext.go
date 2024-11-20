package authorize

import (
	"net/http"
)

const (
	PAGINATION_HAL_LINK_INDEX_SELF = "self"
	PAGINATION_HAL_LINK_INDEX_NEXT = "next"
	PAGINATION_HAL_LINK_INDEX_PREV = "prev"
)


type PagedCursor[T any] struct {
	Data  *T
	HTTPResponse *http.Response
}

func (o *PagedCursor[T]) hasHalLink(linkIndex string) bool {

	if o.Data == nil {
		return false
	}

	if l, ok := o.Data.GetLinksOk(); ok && l != nil {
		links := *l
		if v, ok := links[linkIndex]; ok {
			if h, ok := v.GetHrefOk(); ok && h != nil && *h != "" {
				return true
			}
		}
	}
	return false
}

func (o *PagedCursor[T]) getHalLink(linkIndex string) LinksHATEOASValue {

	var ret LinksHATEOASValue
	if o.Data == nil {
		return ret
	}

	if l, ok := o.Data.GetLinksOk(); ok && l != nil {
		links := *l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	return ret
}

func (o *PagedCursor[T]) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {

	if o.Data == nil {
		return nil, false
	}

	if l, ok := o.Data.GetLinksOk(); ok && l != nil {
		links := *l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o *PagedCursor[T]) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o *PagedCursor[T]) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o *PagedCursor[T]) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o *PagedCursor[T]) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o *PagedCursor[T]) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o *PagedCursor[T]) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o *PagedCursor[T]) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o *PagedCursor[T]) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o *PagedCursor[T]) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o *PagedCursor[T]) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}
