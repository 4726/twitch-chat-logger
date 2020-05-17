import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, FormControl } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.scss']
})
export class SearchComponent implements OnInit {
  public searchForm: FormGroup;
  constructor(
    public formBuilder: FormBuilder,
    private router: Router,
    ){
      this.searchForm = this.formBuilder.group({
        channel: '',
        term: '',
        name: '',
        day: '',
        month: '',
        year: '',
        subscribe_min: '',
        admin: new FormControl(false),
        global_mod: new FormControl(false),
        moderator: new FormControl(false),
        staff: new FormControl(false),
        turbo: new FormControl(false),
        bits_min: '',
        bits_max: '',
      });
    }


  ngOnInit(): void {
  }

  onSubmit(postData) {
    this.router.navigate(['/messages'], {queryParams: postData})
  }
}
