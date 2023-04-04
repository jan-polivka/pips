import { SimulationDisplay } from '@/components/SimulationDisplay';
import { getMatch } from '@/logic/fetching';
import { Button } from '@mui/material'
import Grid from '@mui/material/Unstable_Grid2';
import { useEffect, useState } from 'react';

export default function Home() {


  const [gptMatrix, setGptMatrix] = useState<number[][]>([])
  const [selectedMatrix, setSelectedMatrix] = useState<number[]>([])
  const [idx, setIdx] = useState(0)
  const [isRunning, setIsRunning] = useState(false);


  const onClickGetMatch = (async () => {
    getMatch().then((res: any) => {
      setGptMatrix(JSON.parse(res["message"]))
    })
  })


  useEffect(() => {
    let intervalId: NodeJS.Timer;

    if (isRunning) {
      intervalId = setInterval(() => {
        if (idx < gptMatrix.length) {
          const newSelectedMatrix = gptMatrix[idx];
          setSelectedMatrix(newSelectedMatrix);
          setIdx(idx + 1)
        } else {
          setIsRunning(!isRunning)
          setIdx(0)
          setSelectedMatrix([])
        }
      }, 100);
    }

    return () => clearInterval(intervalId);
  }, [selectedMatrix, isRunning, gptMatrix, idx]);

  const handleClick = () => {
    setIsRunning(!isRunning);
  };

  return (
    <>
      <div>
        <Grid container spacing={2}>
          <Grid xs={12} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <SimulationDisplay matrix={selectedMatrix} />
          </Grid>
          <Grid xs={6} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <Button variant="contained" onClick={() => onClickGetMatch()}>Get match</Button>
          </Grid >
          <Grid xs={6} display={'flex'} justifyContent={'center'} alignContent={'center'}>
            <Button variant="contained" onClick={() => handleClick()}>Play match</Button>
          </Grid>
        </Grid >
      </div >

    </>
  )
}
