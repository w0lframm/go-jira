import { Component, Input } from '@angular/core';
// import { User } from 'src/app/model/cardUser';

@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.css'],
})
export class CardComponent {
  @Input() firstName: string = '';
  @Input() lastName: string = '';
  @Input() avatar: string = '';
  // @Input() cvURL?: string = '';
  @Input() gitURL: string = '';
}
