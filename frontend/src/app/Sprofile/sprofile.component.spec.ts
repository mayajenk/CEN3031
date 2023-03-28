import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SprofileComponent } from './sprofile.component';

describe('SprofileComponent', () => {
  let component: SprofileComponent;
  let fixture: ComponentFixture<SprofileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SprofileComponent ]
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
