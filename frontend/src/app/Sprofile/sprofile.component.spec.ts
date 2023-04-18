import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatChipsModule } from '@angular/material/chips';
import { MatDialogModule } from '@angular/material/dialog';
import { SprofileComponent } from './sprofile.component';
import { MatCardModule } from '@angular/material/card';

describe('SprofileComponent', () => {
  let component: SprofileComponent;
  let fixture: ComponentFixture<SprofileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SprofileComponent ],
      imports: [HttpClientTestingModule, MatChipsModule, MatDialogModule, MatCardModule],
    })
    .compileComponents();

    fixture = TestBed.createComponent(SprofileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
