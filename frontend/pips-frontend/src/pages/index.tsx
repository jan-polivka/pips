import { SimulationDisplay } from '@/components/SimulationDisplay';
import { getMatch } from '@/logic/fetching';
import { Matrix } from '@/store/matrix';
import { Button } from '@mui/material'
import Grid from '@mui/material/Unstable_Grid2';
import { createContext, useEffect, useState } from 'react';

export const MatrixContext = createContext("")

export default function Home() {

  const matrix = new Matrix()

  let intervalId: NodeJS.Timer


  const onClickGetMatch = (async () => {
    getMatch().then((res: any) => {
      matrix.setMatrix(JSON.parse(res["message"]))
    })
  })

  const onClickStartMatch = (() => {
    intervalId = setInterval(() => {
      if (matrix.idx < matrix.matrix.length) {
        matrix.incrementIdx()
        console.log(JSON.stringify(matrix.matrix[matrix.idx]))
      } else {
        clearInterval(intervalId)
        matrix.resetIdx()
      }
    }, 1000)
  })

  return (
    <>
      <div>
        <Grid container spacing={2}>
          <Grid xs={12} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <SimulationDisplay matrix={matrix} />
          </Grid>
          <Grid xs={6} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <Button variant="contained" onClick={() => onClickGetMatch()}>Get match</Button>
          </Grid >
          <Grid xs={6} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <Button variant="contained" onClick={() => onClickStartMatch()}>Play match</Button>
          </Grid>
        </Grid >
      </div >

    </>
  )
}
