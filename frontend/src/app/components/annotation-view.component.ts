// src/app/components/annotation-view.component.ts

import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-annotation-view',
  template: `
    <div>
      <h2>{{ annotation.title }}</h2>
      <p>{{ annotation.content }}</p>
    </div>
  `,
})
export class AnnotationViewComponent {
  @Input() annotation: any;
}
