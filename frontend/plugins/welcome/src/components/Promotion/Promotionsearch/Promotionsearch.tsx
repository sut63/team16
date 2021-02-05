import React, { useState, useEffect } from 'react';
import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
    InfoCard,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { Alert } from '@material-ui/lab';

import { makeStyles, Theme, createStyles, ThemeProvider } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../../api/apis';

import SearchIcon from '@material-ui/icons/Search';
import ArrowBackIcon from '@material-ui/icons/ArrowBack'; // back icon

import { EntPromotion } from '../../../api/models';


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
    promotionsearchcheck: true,
}

export default function Promotionsearch() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [promotion, setPromotion] = useState<EntPromotion[]>([]);
  
    const [promotionsearch, setPromotionsearch] = useState(String);
    
  useEffect(() => {
    const getPromotion = async () => {
      const res = await api.listPromotion({limit:10000, offset:0});
      setLoading(false);
      setPromotion(res);
    };
    getPromotion();
    }, [loading]);

    const SearchPromotion = async () => {
        const res = await api.listPromotion({limit:10000, offset:0});
        const promotionsearch = PromotionSearch(res);
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setPromotion([]);
        if(promotionsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setPromotion(promotionsearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.promotionsearchcheck = true;
        
    }

    const PromotionSearch = (res: any) => {
        const data = res.filter((filter: EntPromotion) => filter.cODE?.includes(promotionsearch))
       // console.log(data.length)
        if (data.length != 0 && promotionsearch != "") {
            return data;
        }
        else {
            searchcheck.promotionsearchcheck = false;
            if(promotionsearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const PromotionSearchhandleChange = (event: any) => {
        setPromotionsearch(event.target.value as string);
    };
    
    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="ค้นหาข้อมูลโปรโมชั่น">

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>รหัสโปรโมชั่น</strong></div>
                    <TextField
                        id="code"
                       // label="ค้นหารหัสพนักงาน"
                        type="string"
                        size="medium"
                        value={promotionsearch}
                        onChange={PromotionSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>

           

                <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchPromotion();
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
                                <TableCell align="center">ID</TableCell>
                                <TableCell align="center">ชื่อโปรโมชั่น</TableCell>
                                <TableCell align="center">รหัสโปรโมชั่น</TableCell>
                                <TableCell align="center">เงื่อนไขโปรโมชั่น</TableCell>
                                <TableCell align="center">วันที่หมดอายุ</TableCell>
                                <TableCell align="center">จำนวนการใช้งาน</TableCell>
                                <TableCell align="center">ชนิดโปรโมชั่น</TableCell>
                                <TableCell align="center">ผู้รับผิดชอบ</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {promotion.map((item:EntPromotion ) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.nAME}</TableCell>
                                    <TableCell align="center">{item.cODE}</TableCell>
                                    <TableCell align="center">{item.dESC}</TableCell>
                                    <TableCell align="center">{item.dATE}</TableCell>
                                    <TableCell align="center">{item.edges?.promotionamount?.aMOUNT}</TableCell>
                                    <TableCell align="center">{item.edges?.promotiontype?.tYPE}</TableCell>
                                    <TableCell align="center">{item.edges?.employee?.eMPLOYEENAME}</TableCell> 
                                </TableRow>
                                ))}
                            </TableBody>
                            </Table>
                        </TableContainer>

                        <div>{status ? (
                          <div>{alerttype != "" ? (
                            <Alert 
                              severity={alerttype} 
                              style={{
                                width: 400 ,
                                marginTop: 20, 
                                marginLeft:6 }} 
                                onClose={() => { 
                                  setStatus(false) 
                                }}
                            >
                              {errormessege}
                            </Alert>
                          ) : null}
                          </div>
                        ) : null}
                        </div>

                </InfoCard>
                <Button
                style={{ width: 110, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                variant="contained"
                color="primary"
                startIcon={<ArrowBackIcon />}
                href="./Promotion"
              >
                ย้อนกลับ
             </Button>
            </Content>
     </Page>
    );
}