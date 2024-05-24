import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Container, Box, Button, Typography, Card, CardContent, Grid, Avatar, Divider } from '@mui/material';
import { styled } from '@mui/system';
import { deepOrange } from '@mui/material/colors';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import ArrowForwardIcon from '@mui/icons-material/ArrowForward';
import LogoutIcon from '@mui/icons-material/Logout';

const RootBox = styled(Box)(({ theme }) => ({
  marginTop: theme.spacing(8),
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
}));

const CustomCard = styled(Card)(({ theme }) => ({
  borderRadius: '12px',
  minWidth: 256,
  textAlign: 'center',
  boxShadow: '0 2px 4px -2px rgba(0,0,0,0.24), 0 4px 24px -2px rgba(0, 0, 0, 0.2)',
}));

const CustomAvatar = styled(Avatar)(({ theme }) => ({
  backgroundColor: deepOrange[500],
  width: 60,
  height: 60,
  margin: 'auto',
}));

const CustomDivider = styled(Divider)(({ theme }) => ({
  margin: theme.spacing(2, 0),
}));

const ButtonGroup = styled(Box)(({ theme }) => ({
  marginTop: theme.spacing(2),
  display: 'flex',
  justifyContent: 'space-between',
  width: '100%',
}));

const LogoutButton = styled(Button)(({ theme }) => ({
  marginTop: theme.spacing(2),
}));

const Home = ({ token, handleLogout }) => {
  const [day, setDay] = useState(1);
  const [reading, setReading] = useState(null);
  const [error, setError] = useState('');
  const API_URL = process.env.REACT_APP_API_URL;

  useEffect(() => {
    fetchReading(day);
  }, [day]);

  const fetchReading = async (day) => {
    try {
      const response = await axios.get(`${API_URL}/readings/${day}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setReading(response.data);
      setError('');
    } catch (err) {
      setError('Reading not found');
      setReading(null);
    }
  };

  const nextDay = () => setDay(day + 1);
  const previousDay = () => day > 1 && setDay(day - 1);

  return (
    <Container component="main" maxWidth="md">
      <RootBox>
        <Typography component="h1" variant="h5">
          Bible Reading Plan
        </Typography>
        {error && <Typography color="error">{error}</Typography>}
        {reading && (
          <>
            <Grid container spacing={2} justifyContent="center">
              <Grid item>
                <CustomCard>
                  <CardContent>
                    <CustomAvatar>{reading.Day}</CustomAvatar>
                    <Typography component="h3" variant="h6" sx={{ fontWeight: 'bold', letterSpacing: '0.5px', marginTop: 1, marginBottom: 0 }}>
                      Day {reading.Day}
                    </Typography>
                  </CardContent>
                </CustomCard>
              </Grid>
              <Grid item>
                <CustomCard>
                  <CardContent>
                    <CustomAvatar>P</CustomAvatar>
                    <Typography component="h3" variant="h6" sx={{ fontWeight: 'bold', letterSpacing: '0.5px', marginTop: 1, marginBottom: 0 }}>
                      Period: {reading.Period}
                    </Typography>
                  </CardContent>
                </CustomCard>
              </Grid>
            </Grid>
            <CustomDivider />
            <Grid container spacing={2} justifyContent="center">
              <Grid item xs={12} sm={4}>
                <CustomCard>
                  <CardContent>
                    <Typography component="h3" variant="h6" sx={{ fontWeight: 'bold', letterSpacing: '0.5px', marginBottom: 1 }}>
                      First Reading
                    </Typography>
                    <Typography variant="body1">{reading.FirstReading}</Typography>
                  </CardContent>
                </CustomCard>
              </Grid>
              <Grid item xs={12} sm={4}>
                <CustomCard>
                  <CardContent>
                    <Typography component="h3" variant="h6" sx={{ fontWeight: 'bold', letterSpacing: '0.5px', marginBottom: 1 }}>
                      Second Reading
                    </Typography>
                    <Typography variant="body1">{reading.SecondReading}</Typography>
                  </CardContent>
                </CustomCard>
              </Grid>
              <Grid item xs={12} sm={4}>
                <CustomCard>
                  <CardContent>
                    <Typography component="h3" variant="h6" sx={{ fontWeight: 'bold', letterSpacing: '0.5px', marginBottom: 1 }}>
                      Third Reading
                    </Typography>
                    <Typography variant="body1">{reading.ThirdReading}</Typography>
                  </CardContent>
                </CustomCard>
              </Grid>
            </Grid>
          </>
        )}
        <ButtonGroup>
          <Button
            variant="contained"
            onClick={previousDay}
            disabled={day === 1}
            startIcon={<ArrowBackIcon />}
          >
            Previous
          </Button>
          <Button
            variant="contained"
            onClick={nextDay}
            endIcon={<ArrowForwardIcon />}
          >
            Next
          </Button>
        </ButtonGroup>
        <LogoutButton
          variant="contained"
          color="secondary"
          onClick={handleLogout}
          startIcon={<LogoutIcon />}
        >
          Logout
        </LogoutButton>
      </RootBox>
    </Container>
  );
};

export default Home;
