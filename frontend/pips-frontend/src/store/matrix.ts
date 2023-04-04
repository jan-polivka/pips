import { action, makeAutoObservable, makeObservable, observable } from "mobx"

export class Matrix {
    matrix: Array<Array<number>> = []
    idx: number = 0


    constructor() {
        makeObservable(this, {
            idx: observable,
            incrementIdx: action,
            resetIdx: action,
            setMatrix: action
        })
    }

    incrementIdx() {
        this.idx++
    }

    resetIdx() {
        this.idx = 0
    }

    setMatrix(newMatrix: Array<Array<number>>) {
        this.matrix = newMatrix
    }
}