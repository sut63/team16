import React, { useState, useEffect } from 'react';
import {
    Content,
    Page,
    pageTheme,
    InfoCard,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
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
import { DefaultApi } from '../../../api/apis';

import SearchIcon from '@material-ui/icons/Search';
import ArrowBackIcon from '@material-ui/icons/ArrowBack'; // back icon

import { EntEquipmentrental } from '../../../api/models';


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
  equipmentrentalsearchcheck: true,
}

export default function Equipmentrentalsearch() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [equipmentrental, setEquipmentrental] = useState<EntEquipmentrental[]>([]);
  
    const [equipmentrentalsearch, setEquipmentrentalsearch] = useState(String);
    
  useEffect(() => {
    const getEquipmentrental = async () => {
      const res = await api.listEquipmentrental({limit:10000, offset:0});
      setLoading(false);
      setEquipmentrental(res);
    };
    getEquipmentrental();
    }, [loading]);

    const SearchEquipmentrental = async () => {
        const res = await api.listEquipmentrental({limit:10000, offset:0});
        const equipmentrentalsearch = EquipmentrentalSearch(res);
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setEquipmentrental([]);
        if(equipmentrentalsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setEquipmentrental(equipmentrentalsearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.equipmentrentalsearchcheck = true;
        
    }

    const EquipmentrentalSearch = (res: any) => {
        const data = res.filter((filter: EntEquipmentrental) => filter.edges?.member?.mEMBERNAME?.includes(equipmentrentalsearch))
       // console.log(data.length)
        if (data.length != 0 && equipmentrentalsearch != "") {
            return data;
        }
        else {
            searchcheck.equipmentrentalsearchcheck = false;
            if(equipmentrentalsearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const EquipmentrentalSearchhandleChange = (event: any) => {
        setEquipmentrentalsearch(event.target.value as string);
    };
    
    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="ระบบค้นหาใบยืมอุปกรณ์กีฬา">

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>ชื่อ User ลูกค้า</strong></div>
                    <TextField
                        id="code"
                       // label="ค้นหารหัสพนักงาน"
                        type="string"
                        size="medium"
                        value={equipmentrentalsearch}
                        onChange={EquipmentrentalSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>

           

                <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchEquipmentrental();
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
                                <TableCell align="center">id Equipmentrental</TableCell>
                                <TableCell align="center">ชื่อผู้ยืม</TableCell>
                                <TableCell align="center">ชื่ออุปกรณ์</TableCell>
                                <TableCell align="center">ชนิดอุปกรณ์</TableCell>
                                <TableCell align="center">จำนวนอุปกรณ์</TableCell>
                                <TableCell align="center">วันที่ยืมอุปกรณ์</TableCell>
                                <TableCell align="center">วันที่ครบกำหนดการยืม</TableCell>
                                <TableCell align="center">ชื่อพนักงาน</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {equipmentrental.map((item: EntEquipmentrental) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.member?.mEMBERNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.equipment?.eQUIPMENTNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.equipmenttype?.eQUIPMENTTYPE}</TableCell>
                                    <TableCell align="center">{item.rENTALAMOUNT}</TableCell>
                                    <TableCell align="center">{item.rENTALDATE}</TableCell>
                                    <TableCell align="center">{item.rETURNDATE}</TableCell>
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
                  href="./Equipmentrental"
                >
                 ย้อนกลับ
                </Button>
            </Content>
     </Page>
    );
}