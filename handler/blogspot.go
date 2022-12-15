package handler

import (
	"encoding/json"
	"fmt"
	"github.com/jutionck/code-run-2022/model"
	"github.com/jutionck/code-run-2022/utils"
	"io"
	"net/http"
)

func (h *Handler) BlogPostList(w http.ResponseWriter, r *http.Request) {
	posts := make([]*model.BlogSpot, 0)
	page := utils.UrlQuery(r, "page")
	size := utils.UrlQuery(r, "size")
	offset, limit := utils.Pagination(page, size)
	err := h.db.Limit(limit).Offset(offset).Find(&posts).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		_, err := fmt.Fprint(w, err)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}

func (h *Handler) BlogPostCreate(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(r.Body)

	var payload model.BlogSpot
	err := json.NewDecoder(r.Body).Decode(&payload)
	// Insert
	err = h.db.Create(&payload).Error
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		_, err := fmt.Fprint(w, err)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(payload)
	if err != nil {
		_, err := fmt.Fprint(w, err)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}

func (h *Handler) BlogPostGet(w http.ResponseWriter, r *http.Request) {
	var post model.BlogSpot
	postID := utils.UrlParams(r, "id")
	err := h.db.First(&post, postID).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		_, err := fmt.Fprint(w, err)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}
