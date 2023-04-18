import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatChipsModule } from '@angular/material/chips';
import { MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { TprofileComponent } from './tprofile.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('TprofileComponent', () => {
  let component: TprofileComponent;
  let fixture: ComponentFixture<TprofileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TprofileComponent ],
      imports: [BrowserAnimationsModule, HttpClientTestingModule, MatChipsModule, MatDialogModule, MatCardModule, MatFormFieldModule],
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
