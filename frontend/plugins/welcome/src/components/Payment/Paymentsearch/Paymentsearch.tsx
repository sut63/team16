import React, { useState, useEffect } from 'react';
import { Header, Content, Page, pageTheme, InfoCard, } from '@backstage/core';
import { TextField, Grid } from '@material-ui/core';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import PowerSettingsNewIcon from '@material-ui/icons/PowerSettingsNew'; // log off icon
import ArrowBackIcon from '@material-ui/icons/ArrowBack'; // back icon

import { DefaultApi } from '../../../api/apis'; // Api Gennerate From Command

import SearchIcon from '@material-ui/icons/Search';

import {
  EntPayment,
} from '../../../api/models';


// header css
const HeaderCustom = {
  minHeight: '20px',
};

// css style 
const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
      margin: theme.spacing(2),
   },
   insideLabel: {
    margin: theme.spacing(1),
  },
   button: {
    marginLeft: '40px',
  },
   textField: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
   },
    select: {
      width: 400 ,
      marginLeft:7,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    center: {
      marginTop: theme.spacing(2),
      marginLeft: theme.spacing(23),
    },
    cardtable: {
      marginTop: theme.spacing(2),
    },
    fieldText: {
      width: 200,
      marginLeft:7,
    },
    fieldLabel: {
      marginLeft:8,
      marginRight: 20,
    },
    table: {
        minWidth: 650,
    }
  }),
);

const searchcheck = {
    paymentsearchcheck: true,
   
}

export default function Paymentsearch() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);
    
    const [payment, setPayment] = React.useState<EntPayment[]>([]);
    const [paymentsearch, setPaymentsearch] = useState(String);
    

  useEffect(() => {
    const getPayment = async () => {
      const res = await api.listPayment({limit:10000, offset:0});
      setLoading(false);
      setPayment(res);
    };
    getPayment();
    }, [loading]);

    const SearchPayment = async () => {
        const res = await api.listPayment({limit:10000, offset:0});
        const paymentsearch = PaymentSearch(res);
        
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setPayment([]);
        if(paymentsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setPayment(paymentsearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.paymentsearchcheck = true;
        
    }
    const PaymentSearch = (res: any) => {
        const data = res.filter((filter:EntPayment) => filter.edges?.member?.mEMBERNAME?.includes(paymentsearch))
       // console.log(data.length)
        if (data.length != 0 && paymentsearch != "") {
            return data;
        }
        else {
            searchcheck.paymentsearchcheck = false;
            if(paymentsearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const PaymentSearchhandleChange = (event: any) => {
        setPaymentsearch(event.target.value as string);
    };
    

    return (
        <Page theme={pageTheme.service}>
            <Content>
            <Header style={HeaderCustom} title={`ค้นหาข้อมูลการชำระเงินของสมาชิก`}>
           </Header>
          
            <InfoCard title="">

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>Member Name</strong></div>
                    <TextField
                        id="memberid"
                       // label="ค้นหารหัสพนักงาน"
                        type="string"
                        size="medium"
                        value={paymentsearch}
                        onChange={PaymentSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>

                    <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchPayment();
                }}
                variant="contained"
                color="primary"
                startIcon={<SearchIcon />}
              >
                ค้นหา
             </Button>
             
             <div><br></br></div>
             <TableContainer component={Paper}>
                            <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                <TableCell align="center">id Payment</TableCell>
                                <TableCell align="center">ชื่อ Member</TableCell>
                                <TableCell align="center">ชื่อ Employee</TableCell>
                                <TableCell align="center">ชื่อ Promotion</TableCell>
                                <TableCell align="center">ประเภทการชำระเงิน</TableCell>
                                <TableCell align="center">ราคา</TableCell>
                                <TableCell align="center">เบอร์ติดต่อ</TableCell>
                                <TableCell align="center">Email</TableCell>
                                <TableCell align="center">วันที่ชำระเงิน</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {payment.map((item:EntPayment ) => (
                                  <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.member?.mEMBERNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.employee?.eMPLOYEENAME}</TableCell>
                                    <TableCell align="center">{item.edges?.promotion?.nAME}</TableCell>
                                    <TableCell align="center">{item.edges?.paymenttype?.tYPE}</TableCell>
                                    <TableCell align="center">{item.pAYMENTAMOUNT}</TableCell>
                                    <TableCell align="center">{item.pHONENUMBER}</TableCell>
                                    <TableCell align="center">{item.eMAIL}</TableCell> 
                                    <TableCell align="center">{item.pAYMENTDATE}</TableCell> 
                                </TableRow>
                                ))}
                            </TableBody>
                            </Table>
                        </TableContainer>

                        <div>
             {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} style={{ width: 400 ,marginTop: 20, marginLeft:6 }} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}</div>

                </InfoCard>
                <Button
                  style={{ width: 110, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                  variant="contained"
                  color="primary"
                  startIcon={<ArrowBackIcon />}
                  href="./Payment"
                >
                 ย้อนกลับ
                </Button>

            </Content>
     </Page>
    );
}