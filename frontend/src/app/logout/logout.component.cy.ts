import { HttpClientModule } from "@angular/common/http"
import { MatButtonModule } from "@angular/material/button"
import { MatDialogModule, MatDialogRef } from "@angular/material/dialog"
import { AuthService } from "../auth/auth.service"
import { LogoutComponent } from "./logout.component"

describe('LogoutComponent', () => {
    it('mounts', () => {
        cy.mount(LogoutComponent, {
            imports: [MatDialogModule, HttpClientModule, MatButtonModule],
            providers: [AuthService, {provide : MatDialogRef, useValue : {}}]
        })
    });
});