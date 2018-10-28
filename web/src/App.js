import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {
  withStyles,
  MuiThemeProvider,
  createMuiTheme
} from '@material-ui/core/styles';
import {
  AppBar,
  Toolbar,
  Paper,
  CssBaseline,
  Typography
} from '@material-ui/core';
import { red, yellow } from '@material-ui/core/colors/red';
import { withFetching, TemperaturesChart } from './components';

const styles = theme => ({
  root: {
    flexGrow: 1
  },
  content: {
    margin: theme.spacing.unit * 2,
    padding: theme.spacing.unit * 2,
    color: theme.palette.text.secondary
  },
  chart: {
    height: '800px',
    backgroundColor: '#fff',
    marginTop: theme.spacing.unit * 2
  }
});

const theme = createMuiTheme({
  typography: {
    useNextVariants: true
  },
  palette: {
    type: 'dark',
    primary: red,
    secondary: yellow
  }
});

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { chartData: null };
  }

  componentDidUpdate() {
    const { data } = this.props;

    if (!data || this.state.chartData) {
      return;
    }
    const chartData = data.thermostats.map(t => ({
      id: t.name,
      data: t.temperatures.map(temp => ({
        x: temp.timestamp.slice(0, 19), // strip milliseconds and timezone
        y: temp.ambientTemperatureC
      }))
    }));
    this.setState(prevState => ({ ...prevState, chartData }));
  }

  render() {
    const { classes, error, isLoading } = this.props;
    const { chartData } = this.state;

    return (
      <MuiThemeProvider theme={theme}>
        <CssBaseline />
        <div className={classes.root}>
          <AppBar position="static" color="default">
            <Toolbar>
              <Typography variant="h6" color="inherit">
                Go thermostat
              </Typography>
            </Toolbar>
          </AppBar>
          <Paper className={classes.content} color="default">
            <Typography variant="h6">Go thermostat</Typography>
            <Typography variant="body1">
              Control your nest thermostat and view your temperature stats.
            </Typography>
            <Paper className={classes.chart}>
              {!chartData || error || isLoading ? (
                'no data'
              ) : (
                <TemperaturesChart data={chartData} />
              )}
            </Paper>
          </Paper>
        </div>
      </MuiThemeProvider>
    );
  }
}

App.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired,
  isLoading: PropTypes.bool.isRequired,
  data: PropTypes.object,
  error: PropTypes.object
};

export default withFetching('/api/thermostat-data')(
  withStyles(styles, { withTheme: true })(App)
);
