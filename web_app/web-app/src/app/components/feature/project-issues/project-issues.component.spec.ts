import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProjectIssuesComponent } from './project-issues.component';

describe('ProjectIssuesComponent', () => {
  let component: ProjectIssuesComponent;
  let fixture: ComponentFixture<ProjectIssuesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProjectIssuesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProjectIssuesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
