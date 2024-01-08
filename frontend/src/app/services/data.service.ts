// src/app/services/data.service.ts

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class DataService {
  private apiUrl = 'http://localhost:8080/api'; // Adjust the URL based on your backend's location

  constructor(private http: HttpClient) {}

  getBooks(): Observable<any> {
    return this.http.get(`${this.apiUrl}/books`);
  }
}
