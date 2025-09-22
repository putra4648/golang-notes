export type Note = {
    readonly id: string;
    readonly title: string;
    readonly content: string;
    readonly created?: Date;
    readonly updated?: Date;
    readonly isDone: boolean
}

