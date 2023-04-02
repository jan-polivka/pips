import { getMatch } from '@/logic/fetching';
import { Button } from '@mui/material'
import Grid from '@mui/material/Unstable_Grid2';
import { Inter } from 'next/font/google'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
      <div>
        <Grid container spacing={2}>
          <Grid xs={12} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            Hello
          </Grid>
          <Grid xs={6} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <Button variant="contained" onClick={() => getMatch()}>Get match</Button>
          </Grid >
          <Grid xs={6} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <Button variant="contained">Play match</Button>
          </Grid>
        </Grid >
      </div >

    </>
  )
}
