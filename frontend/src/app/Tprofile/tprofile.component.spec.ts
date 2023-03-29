import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatChipsModule } from '@angular/material/chips';

import { TprofileComponent } from './tprofile.component';

describe('TprofileComponent', () => {
  let component: TprofileComponent;
  let fixture: ComponentFixture<TprofileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TprofileComponent ],
      imports: [HttpClientTestingModule, MatChipsModule],
    })
    .compileComponents();

    fixture = TestBed.createComponent(TprofileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
