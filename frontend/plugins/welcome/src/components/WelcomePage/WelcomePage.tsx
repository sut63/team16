import React, { FC } from 'react';
import { Typography, Grid, Link, Button } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import { Cookies } from '../../cookies';
import PowerSettingsNewIcon from '@material-ui/icons/PowerSettingsNew'; //log off icon
import Avatar from '@material-ui/core/Avatar';

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles((theme) => ({
  root: {
      maxWidth: 345,
      display: 'flex',
      '& > *': {
        margin: theme.spacing(1),
      },
    },
}));

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
  linkto: string;
};

export function CardTeam({ name, id, system, linkto }: ProfileProps) {
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
          <Link 
            href={linkto}
          > 
          <CardMedia
            component="img"
            //alt="นาย สมชาย ใจดี"
            height="140"
            //image="../../image/promotionimage.png"
            //title="นาย สมชาย ใจดี"
          />
          </Link>
          <CardContent>
            <Typography gutterBottom variant="h5" component="h2">
              {system}
            </Typography>
            <Typography gutterBottom variant="h5" component="h2">
              {id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const WelcomePage: FC<{}> = () => {
  const classes = useStyles();
  var ck = new Cookies()
  var cookieName = ck.GetCookie()

  // clear cookie
  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบสถานกีฬา`}>{cookieName}
      <div className={classes.root}>
        <Avatar src="/broken-image.jpg" />
      </div>
       <Grid item xs={9}></Grid>
        <Grid item xs>
          <Button
            variant="contained"
            color="secondary"
            size="large"
            startIcon={<PowerSettingsNewIcon />}
            href="/"
            onClick={Clears}
          >
            sign out
          </Button>
        </Grid>
      </Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Grid container>
          <CardTeam name={"นายวิศรุต เทศารินทร์"} id={"B6011017"} system={"ระบบชำระเงิน"} linkto="/Payment" ></CardTeam>
          <CardTeam name={"นายสุธี สีสิงห์"} id={"B6025427"} system={"ระบบเพิ่มอุปกรณ์"} linkto="/Equipment" ></CardTeam>
          <CardTeam name={"นาย ชยากร พิลึกนา"} id={"B6005818"} system={"ระบบโปรโมชั่น"} linkto="/Promotion" ></CardTeam>
          <CardTeam name={"นาย กิตติศักดิ์ เพชรแหน"} id={"B6015909"} system={"ระบบจองสนามกีฬา"} linkto="/Bookcourse" ></CardTeam>
          <CardTeam name={"นาย ชนุดม ศรีสุรัตน์"} id={"B6006297"} system={"ระบบเพิ่มข้อมูลพนักงาน"} linkto="/Employee" ></CardTeam>
          <CardTeam name={"นางสาว ศิริลักษณ์ รองกระโทก"} id={"B6010256"} system={"ระบบยืมอุปกรณ์กีฬา"} linkto="/Equipmentrental" ></CardTeam>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
