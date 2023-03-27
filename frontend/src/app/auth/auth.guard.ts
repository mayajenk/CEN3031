import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, UrlTree, Router } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  $isLoggedIn: Observable<boolean>;

  constructor(private authService: AuthService, private router: Router){
    this.$isLoggedIn = this.authService.isLoggedIn$;
  };

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): boolean {
    
    if (this.$isLoggedIn) {
      return true;
    }
    else {
      this.router.navigate(['/login']);
      return false;
    }
  }
}
