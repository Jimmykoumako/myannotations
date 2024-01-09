// src/app/components/book-list.component.ts

import { Component, OnInit } from '@angular/core';
import { DataService } from '../services/data.service';

@Component({
  selector: 'app-book-list',
  template: `
    <h2>Book List</h2>
    <ul>
      <li *ngFor="let book of books">{{ book.title }}</li>
    </ul>
  `,
})
export class BookListComponent implements OnInit {
  books: any[] = [];

  constructor(private dataService: DataService) {}

  ngOnInit() {
    this.dataService.getBooks().subscribe((data) => {
      this.books = data;
    });
  }
}
