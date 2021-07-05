import { Component } from '@angular/core';
import { CardModel } from './models/card.model';
import { DataService } from './services/data.service';
import {webSocket} from 'rxjs/webSocket';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'deliveries';
  cards: CardModel[];
  subject = webSocket('ws://localhost:8080/');

  constructor( dataService: DataService ){
    this.cards = dataService.getAllRoutes();
    this.subject.subscribe(    
      msg => {
        const found = this.cards.find(element => element.ruta == msg['route_id']);
        found.entregas = msg['deliveries']
        found.completada = msg['status'] == 'completed' ? true : false;
        found.termino = msg['completed_at']!="" ? msg['completed_at'].split('T')[1].split('-')[0] : '-'
      }, 
      // Called whenever there is a message from the server    
      err => console.log(err), 
      // Called if WebSocket API signals some kind of error    
      () => console.log('complete') 
      // Called when connection is closed (for whatever reason)  
   );
  }
}
