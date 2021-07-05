import { Injectable } from '@angular/core';
import { CardModel } from '../models/card.model';
import data from './routes.json';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  constructor() { }

  getAllRoutes(): CardModel[] {
    let cards = [];
    for (const item of data.routes) {
      let card: CardModel = {
        conductor: item['driver_name'],
        ruta : item['id'].toString(),
        creada : item['created_at'].split('T')[0],
        inicio : item['created_at'].split('T')[1].split('-')[0],
        entregas : item['deliveries'],
        completada : false,
        termino : '-'
      };
      cards.push(card)
    }
    return cards;
  }
}
