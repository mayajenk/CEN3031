import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { StudentDialogComponent } from './student-dialog.component';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule, MatFormFieldControl } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('StudentDialogComponent', () => {
  let component: StudentDialogComponent;
  let fixture: ComponentFixture<StudentDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ StudentDialogComponent ],
      imports: [HttpClientModule, FormsModule, MatFormFieldModule, MatInputModule, BrowserAnimationsModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StudentDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
