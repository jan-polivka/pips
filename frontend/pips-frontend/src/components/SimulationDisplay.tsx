import { matrixReducer } from "@/logic/matrixReducer";
import { MatrixContext } from "@/pages";
import { useContext, useEffect, useReducer, useState } from "react";

export var SimulationDisplay = ((props: { matrixProp: Array<Number> }): JSX.Element => {

    const matrix = useContext(MatrixContext)


    //Now, just setup the timer again and on you go with the Matrix?

    return (
        <>
            {matrix}
        </>
    )


})