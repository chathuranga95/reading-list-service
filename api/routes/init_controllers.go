// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package routes

import (
	"reading-list-service/internal/config"
	"reading-list-service/internal/controllers"
	"reading-list-service/internal/repositories"
)

var bookController *controllers.BookController

func initControllers() {
	initialData := config.LoadInitialData()
	bookRepository := repositories.NewBookRepository(initialData.Books)
	bookController = controllers.NewBookController(bookRepository)
}
