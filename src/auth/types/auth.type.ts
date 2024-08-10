import { ApiProperty } from "@nestjs/swagger"

export class AuthType{
    @ApiProperty()
    access_token:string

    @ApiProperty()
    refresh_token:string
}