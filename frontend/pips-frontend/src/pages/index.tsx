import { SimulationDisplay } from '@/components/SimulationDisplay';
import { getMatch } from '@/logic/fetching';
import { Button } from '@mui/material'
import Grid from '@mui/material/Unstable_Grid2';
import { Inter } from 'next/font/google'
import { createContext, useState } from 'react';

const inter = Inter({ subsets: ['latin'] })

export const MatrixContext = createContext("")

export default function Home() {

  const [matrix, setMatrix] = useState<string>("")

  const onClickGetMatch = (async () => {
    getMatch().then((res: any) => {
      setMatrix(res["message"])
    })
  })

  return (
    <>
      <div>
        <Grid container spacing={2}>
          <Grid xs={12} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <MatrixContext.Provider value={matrix}>
              <SimulationDisplay matrixProp={[1, 2, 3]} />
            </MatrixContext.Provider>
          </Grid>
          <Grid xs={6} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <Button variant="contained" onClick={() => onClickGetMatch()}>Get match</Button>
          </Grid >
          <Grid xs={6} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <Button variant="contained">Play match</Button>
          </Grid>
        </Grid >
      </div >

    </>
  )
}
