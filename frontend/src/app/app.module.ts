import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClient, HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { CookieService } from 'ngx-cookie-service';

import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { PathLocationStrategy, LocationStrategy } from '@angular/common';
import { AuthService } from './auth/auth.service'
import { ProfileComponent } from './profile/profile.component';
import { SearchComponent } from './search/search.component';
import { LogoutComponent } from './logout/logout.component';
import { SprofileComponent } from './Sprofile/sprofile.component';
import { TprofileComponent } from './Tprofile/tprofile.component';
import { AuthGuard } from './auth/auth.guard';

import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatChipsModule } from '@angular/material/chips';
import { MatListModule } from '@angular/material/list';
import { MatDividerModule } from '@angular/material/divider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';
import {MatToolbarModule} from '@angular/material/toolbar'; 
import {MatDialogModule} from '@angular/material/dialog';
import { MatSelectModule } from '@angular/material/select';
import { HomeComponent } from './home/home.component';
import { MatLabel } from '@angular/material/form-field';
import { MatFormField } from '@angular/material/form-field';
import { DialogComponent } from './dialog/dialog.component';
import { SearchProfileComponent } from './search-profile/search-profile.component';
import { StudentDialogComponent } from './student-dialog/student-dialog.component';
 

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    RegisterComponent,
    ProfileComponent,
    SearchComponent,
    LogoutComponent,
    SprofileComponent,
    TprofileComponent,
    DialogComponent,
    SearchProfileComponent,
    StudentDialogComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MatButtonToggleModule,
    MatChipsModule,
    MatDividerModule,
    MatListModule,
    MatFormFieldModule,
    MatIconModule,
    MatButtonModule,
    MatInputModule,
    MatCardModule,
    MatToolbarModule,
    MatDialogModule,
    MatSelectModule,
    FormsModule
  ],
  providers: [HttpClient, CookieService, AuthService, AuthGuard, {provide: LocationStrategy, useClass: PathLocationStrategy}],
  bootstrap: [AppComponent]
})
export class AppModule { }
