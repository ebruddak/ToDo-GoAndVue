import {Entity} from "@/interfaces/entity";

export class Todo implements Entity {
    id: number;
    title: string;
    content: string;
    userId: string;
    priority: string;
    groupId: string;


    constructor(id = 0,title='', content = '', userId = '',priority='',groupId='') {
        this.id = id;
        this.title = title;
        this.userId = userId;
        this.priority=priority;
        this.groupId=groupId;
        this.content=content;
    }
}