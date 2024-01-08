// src/app/app.module.ts

import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { AnnotationViewComponent } from './components/annotation-view.component';
import { BookListComponent } from './components/book-list.component';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [AppComponent, AnnotationViewComponent, BookListComponent],
  imports: [BrowserModule, HttpClientModule],
  bootstrap: [AppComponent],
})
export class AppModule {}
