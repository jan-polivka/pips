import { matrixReducer } from "@/logic/matrixReducer";
import { MatrixContext } from "@/pages";
import { observer } from "mobx-react-lite";
import { useContext, useEffect, useReducer, useState } from "react";

export var SimulationDisplay = observer((({ matrix }: any): JSX.Element => {


    return (
        <>
            {JSON.stringify(matrix.matrix[matrix.idx])}
        </>
    )


}))
