import {Entity} from "@/interfaces/entity";

export class Group implements Entity {
    id: number;
    name: string;
    userId: string;
   

    constructor(id = 0, name = '', userId = '') {
        this.id = id;
        this.name = name;
        this.userId = userId;

    }
}