import { matrixReducer } from "@/logic/matrixReducer";
import { useEffect, useReducer, useState } from "react";

export var SimulationDisplay = ((props: { matrixProp: Array<Number> }): JSX.Element => {


    const [time, setTime] = useState(Date.now());
    const [idx, setIdx] = useState(0)
    const [matrix, setMatrix] = useReducer(matrixReducer, props.matrixProp)

    useEffect(() => {
        const interval = setInterval(() => {
            setTime(Date.now())
            if (idx < matrix.length) {
                console.log(matrix[idx])
            }
            console.log(idx + 1)
            setIdx(idx + 1)
        }, 1000);
        return () => {
            console.log("clearing")
            clearInterval(interval);
        };
    }, []);

    return (
        <>
            {time}
            <br></br>
            {idx}
        </>
    )


})