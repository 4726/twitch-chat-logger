import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.scss']
})
export class SearchComponent implements OnInit {
  private searchForm: FormGroup;
  private router: Router;
  constructor(
    private formBuilder: FormBuilder,
    ){
      this.searchForm = this.formBuilder.group({
        channel: '',
        term: '',
        name: '',
        date: '',
        subscribe_min: 0,
        admin: false,
        global_mod: false,
        staff: false,
        turbo: false,
        bits_min: 0,
        bits_max: 0,
      });
    }


  ngOnInit(): void {
  }

  onSubmit(postData) {
    this.router.navigate(['/messages'], {queryParams: postData})
  }
}
