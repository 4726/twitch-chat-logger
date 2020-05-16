import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.scss']
})
export class MessagesComponent implements OnInit {
  public messages: Message[];
  constructor() { 
  }

  ngOnInit(): void {
    this.testInit()
  }

  testInit(): void {
    let msg: Message = new Message();
    msg.Text = "hello world"
    msg.DisplayName = "user1"
    let msg2: Message = new Message();
    msg2.Text = "hello"
    msg2.DisplayName = "user2"
    this.messages = [
      msg, msg2
    ]
  }

}

class Message {
  Text: string;
  DisplayName: string;
}