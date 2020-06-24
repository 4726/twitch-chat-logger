import { Component, OnInit } from '@angular/core';
import {environment} from '../../environments/environment';
import { HttpClient, HttpParams, HttpClientModule } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.scss']
})
export class MessagesComponent implements OnInit {
  public messages : Message[];
  private backendParams;
  constructor(private http: HttpClient, 
    private route: ActivatedRoute
  ) { 
    this.route.queryParams.subscribe(params => {
      this.backendParams = {
        Channel:  params['channel'],
        Name:  params['name'],
        Date:  params['day'] + params['month'] + params['year'],
      }
    })
  }

  ngOnInit(): void {
    this.callAPI()
  }

  callAPI() {
    this.http.get<BackendMessage>(environment.backend + '/messages/search', {
      params: this.backendParams
    })
      .subscribe(
        res => {
          this.messages = res.messages
        }
      )
  }
}

class BackendMessage {
  messages: Message[];
}

class Message {
  Message: string;
  DisplayName: string;
}

class BackendParams {
	Channel:      string
	Term:         string
	Name:         string
	Date:         string
	SubscribeMin: number
	Admin:        boolean
	GlobalMod:    boolean
	Staff:        boolean
	Turbo:        boolean
	BitsMin:      number
	BitsMax:      number
}
