import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { HeptagonGridComponent } from '../heptagon-grid/heptagon-grid.component';
import { PodStatusComponent } from './pod-status.component';
import { Input, Component } from '@angular/core';
import { PodStatus } from '../../models/pod-status';

@Component({
  selector: 'app-heptagon-grid',
  template: ``
})
class TestGridComponent {
  @Input()
  podStatuses: PodStatus[] = [];

  @Input()
  edgeLength: number;
}

describe('PodStatusComponent', () => {
  let component: PodStatusComponent;
  let fixture: ComponentFixture<PodStatusComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [PodStatusComponent, TestGridComponent],
      providers: [
        { provide: HeptagonGridComponent, useClass: TestGridComponent },
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PodStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
