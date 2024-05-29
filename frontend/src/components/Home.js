import React, { useState, useEffect } from 'react';
import axios from 'axios';
import {
  Container, Box, Button, Typography, Card, CardContent, Grid, Avatar, Divider,
  BottomNavigation, BottomNavigationAction, FormControlLabel, Switch, AppBar, Toolbar, Menu, MenuItem, IconButton
} from '@mui/material';
import { styled } from '@mui/system';
import { deepOrange, blue, grey } from '@mui/material/colors';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import ArrowForwardIcon from '@mui/icons-material/ArrowForward';
import LogoutIcon from '@mui/icons-material/Logout';
import MenuIcon from '@mui/icons-material/Menu';

const API_URL = process.env.REACT_APP_API_URL;

const RootBox = styled(Box)(({ theme }) => ({
  marginTop: theme.spacing(8),
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
}));

const CustomCard = styled(Card)(({ theme }) => ({
  borderRadius: '12px',
  minWidth: 256,
  width: '100%',
  textAlign: 'center',
  boxShadow: '0 2px 4px -2px rgba(0,0,0,0.24), 0 4px 24px -2px rgba(0, 0, 0, 0.2)',
  maxHeight: '400px',
  overflowY: 'auto'
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

const StatusGrid = styled(Box)(({ theme }) => ({
  display: 'grid',
  gridTemplateColumns: 'repeat(40, 2fr)',
  gap: theme.spacing(0.5),
  marginTop: theme.spacing(2),
}));

const DaySquare = styled(Box)(({ theme, isRead }) => ({
  width: theme.spacing(3),
  height: theme.spacing(3),
  backgroundColor: isRead ? blue[500] : grey[300],
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  color: '#fff',
  borderRadius: theme.shape.borderRadius,
  cursor: 'pointer',
  fontSize: '0.75 rem',
}));

const Home = ({ token, handleLogout }) => {
  const [day, setDay] = useState(1);
  const [reading, setReading] = useState(null);
  const [status, setStatus] = useState({});
  const [selectedReading, setSelectedReading] = useState('FirstReading');
  const [anchorEl, setAnchorEl] = useState(null);
  const [readingText, setReadingText] = useState('');
  const userId = localStorage.getItem('userId');

  useEffect(() => {
    const fetchReading = async () => {
      try {
        const response = await axios.get(`${API_URL}/readings/${day}`);
        setReading(response.data);
        setSelectedReading('FirstReading');
      } catch (error) {
        console.error("Failed to fetch reading:", error);
      }
    };

    const fetchStatus = async () => {
      try {
        const response = await axios.get(`${API_URL}/status/${userId}`);
        const statusMap = response.data.reduce((acc, item) => {
          acc[item.day] = item.status;
          return acc;
        }, {});
        setStatus(statusMap);
      } catch (error) {
        console.error("Failed to fetch status:", error);
      }
    };

    fetchReading();
    fetchStatus();
  }, [day, userId]);

  useEffect(() => {
    if (reading) {
      const fetchReadingText = async () => {
        try {
          const description = reading[selectedReading];
          if (!description) {
            setReadingText("Não há segunda leitura para o dia de hoje");
            return;
          }
          const response = await axios.get(`${API_URL}/readingText`, {
            params: { description }
          });
          console.log("Reading text response:", response.data);
          setReadingText(response.data.text.join(' '));
        } catch (error) {
          console.error("Failed to fetch reading text:", error);
        }
      };

      fetchReadingText();
    }
  }, [reading, selectedReading]);

  const handleMarkAsRead = async (event) => {
    const newStatus = event.target.checked ? 'read' : 'unread';
    try {
      await axios.post(`${API_URL}/status/${userId}/${day}`, { status: newStatus });
      setStatus((prevStatus) => ({ ...prevStatus, [day]: newStatus }));
    } catch (error) {
      console.error("Failed to update status:", error);
    }
  };

  const handleNext = () => {
    setDay((prevDay) => prevDay + 1);
    setReading(null);
    setReadingText('');
  };

  const handlePrevious = () => {
    if (day > 1) {
      setDay((prevDay) => prevDay - 1);
      setReading(null);
      setReadingText('');
    }
  };

  const isRead = status[day] === 'read';

  const handleMenuOpen = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1, textAlign: 'left' }}>
            Plano de Leitura da Bíblia
          </Typography>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="menu"
            onClick={handleMenuOpen}
          >
            <MenuIcon />
          </IconButton>
          <Menu
            id="menu-appbar"
            anchorEl={anchorEl}
            anchorOrigin={{
              vertical: 'top',
              horizontal: 'right',
            }}
            keepMounted
            transformOrigin={{
              vertical: 'top',
              horizontal: 'right',
            }}
            open={Boolean(anchorEl)}
            onClose={handleMenuClose}
          >
            <MenuItem onClick={handleMenuClose}>Profile</MenuItem>
            <MenuItem onClick={handleMenuClose}>My account</MenuItem>
          </Menu>
          <Button color="inherit" onClick={handleLogout} startIcon={<LogoutIcon />}>
            LogOut
          </Button>
        </Toolbar>
      </AppBar>
      <Container component="main" maxWidth="md">
        <RootBox>
          {reading && (
            <>
              <Typography variant="h6" component="h2" gutterBottom>
                Você está na leitura do {day}º dia.
              </Typography>
              <StatusGrid>
                {Array.from({ length: 365 }, (_, i) => (
                  <DaySquare
                    key={i + 1}
                    isRead={status[i + 1] === 'read'}
                    onClick={() => setDay(i + 1)}
                  >
                    {i + 1}
                  </DaySquare>
                ))}
              </StatusGrid>
              <Typography variant="body1" sx={{ mt: 2 }}>
                Já leu as leituras de hoje?
              </Typography>
              <FormControlLabel
                control={
                  <Switch
                    checked={isRead}
                    onChange={handleMarkAsRead}
                    name="readStatus"
                    color="primary"
                  />
                }
                label="SIM"
                sx={{ mt: 1 }}
              />
              <BottomNavigation
                value={selectedReading}
                onChange={(event, newValue) => setSelectedReading(newValue)}
                showLabels
                sx={{ marginBottom: 2 }}
              >
                <BottomNavigationAction label="Primeira Leitura" value="FirstReading" />
                <BottomNavigationAction label="Segunda Leitura" value="SecondReading" />
                <BottomNavigationAction label="Terceira Leitura" value="ThirdReading" />
              </BottomNavigation>
              <Grid container spacing={2} justifyContent="center">
                <Grid item xs={12}>
                  <CustomCard>
                    <CardContent>
                      <Typography component="h3" variant="h6" sx={{ fontWeight: 'bold', letterSpacing: '0.5px', marginBottom: 1 }}>
                        {selectedReading === 'FirstReading' && 'Primeira Leitura'}
                        {selectedReading === 'SecondReading' && 'Segunda Leitura'}
                        {selectedReading === 'ThirdReading' && 'Terceira Leitura'}
                      </Typography>
                      <Typography variant="body2" color="textSecondary" gutterBottom>
                        {reading[selectedReading]}
                      </Typography>
                      <Typography variant="body1">{readingText}</Typography>
                    </CardContent>
                  </CustomCard>
                </Grid>
              </Grid>
              <CustomDivider />
            </>
          )}
          <ButtonGroup>
            <Button
              variant="contained"
              onClick={handlePrevious}
              disabled={day === 1}
              startIcon={<ArrowBackIcon />}
            >
              Anterior
            </Button>
            <Button
              variant="contained"
              onClick={handleNext}
              endIcon={<ArrowForwardIcon />}
            >
              Próximo
            </Button>
          </ButtonGroup>
        </RootBox>
      </Container>
    </>
  );
};

export default Home;
