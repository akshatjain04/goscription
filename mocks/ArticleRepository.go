// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/kecci/goscription/models"
	mock "github.com/stretchr/testify/mock"
)

// ArticleRepository is an autogenerated mock type for the ArticleRepository type
type ArticleRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ArticleRepository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *ArticleRepository) Fetch(ctx context.Context, cursor string, num int64) ([]models.Article, string, error) {
	ret := _m.Called(ctx, cursor, num)

	var r0 []models.Article
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []models.Article); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Article)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ArticleRepository) GetByID(ctx context.Context, id int64) (models.Article, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Article
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.Article); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *ArticleRepository) GetByTitle(ctx context.Context, title string) (models.Article, error) {
	ret := _m.Called(ctx, title)

	var r0 models.Article
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Article); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(models.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, a
func (_m *ArticleRepository) Store(ctx context.Context, a *models.Article) error {
	ret := _m.Called(ctx, a)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Article) error); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, article
func (_m *ArticleRepository) Update(ctx context.Context, article *models.Article) error {
	ret := _m.Called(ctx, article)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
