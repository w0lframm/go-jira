import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './components/layout/home/home.component';
import { AboutComponent } from './components/layout/about/about.component';
import { ProjectListComponent } from './components/feature/project-list/project-list.component';
import { ProjectIssuesComponent } from './components/feature/project-issues/project-issues.component';
import { ProjectCompareComponent } from './components/feature/project-compare/project-compare.component';

const ROUTES: Routes = [
  { path: '', component: HomeComponent },
  { path: 'about', component: AboutComponent },
  { path: 'projects', component: ProjectListComponent },
  { path: 'issues', component: ProjectIssuesComponent },
  { path: 'compare', component: ProjectCompareComponent },
  { path: '**', component: HomeComponent },
];
@NgModule({
  imports: [RouterModule.forRoot(ROUTES)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
