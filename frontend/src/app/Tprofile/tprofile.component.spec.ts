import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TprofileComponent } from './tprofile.component';

describe('TprofileComponent', () => {
  let component: TprofileComponent;
  let fixture: ComponentFixture<TprofileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TprofileComponent ]
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
