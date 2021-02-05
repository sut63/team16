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

import { EntEquipment } from '../../../api/models';


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
    equipmentsearchcheck: true,
}

export default function Equipmentsearch() {
    const classes = useStyles();
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [loading, setLoading] = useState(true);
    const [alerttype, setAlertType] = useState(String);
    const [errormessege, setErrorMessege] = useState(String);

    const [equipment, setEquipment] = useState<EntEquipment[]>([]);
  
    const [equipmentsearch, setEquipmentsearch] = useState(String);
    
  useEffect(() => {
    const getEquipment = async () => {
      const res = await api.listEquipment({limit:10000, offset:0});
      setLoading(false);
      setEquipment(res);
    };
    getEquipment();
    }, [loading]);

    const SearchEquipment = async () => {
        const res = await api.listEquipment({limit:10000, offset:0});
        const equipmentsearch = EquipmentSearch(res);
        
        setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setEquipment([]);
        if(equipmentsearch.length > 0){
            Object.entries(searchcheck).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setEquipment(equipmentsearch);
                }
            })
        }

        setStatus(true);
        ResetSearchCheck();
    }

    const ResetSearchCheck = () => {
        searchcheck.equipmentsearchcheck = true;
        
    }

    const EquipmentSearch = (res: any) => {
        const data = res.filter((filter: EntEquipment) => filter.eQUIPMENTNAME?.includes(equipmentsearch))
       // console.log(data.length)
        if (data.length != 0 && equipmentsearch != "") {
            return data;
        }
        else {
            searchcheck.equipmentsearchcheck = false;
            if(equipmentsearch == ""){
                return res;
            }
            else{
                return data;
            }
        }
    }

    const EquipmentSearchhandleChange = (event: any) => {
        setEquipmentsearch(event.target.value as string);
    };
    
    return (
        <Page theme={pageTheme.service}>
            <Content>
            <InfoCard title="ค้นหาข้อมูลอุปกรณ์">

            <FormControl
                    className={classes.margin}
                    variant="standard"
                >
                    <div className={classes.paper}><strong>ชื่ออุปกรณ์</strong></div>
                    <TextField
                        id="equipmentname"
                        type="string"
                        size="medium"
                        value={equipmentsearch}
                        onChange={EquipmentSearchhandleChange}
                        style={{ width: 210,marginLeft: 8}}
                    />
                </FormControl>

           

                <Button
                style={{ width: 100, backgroundColor: "#5319e7",marginTop: 49,marginLeft: 20}}
                onClick={() => {
                  SearchEquipment();
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
                                <TableCell align="center">ID Equipment</TableCell>
                                <TableCell align="center">ชื่อของอุปกรณ์</TableCell>
                                <TableCell align="center">ประเภทของอุปกรณ์</TableCell>
                                <TableCell align="center">จำนวนของอุปกรณ์</TableCell>
                                <TableCell align="center">ลักษณะนามของอุปกรณ์</TableCell>
                                <TableCell align="center">ตำแหน่งการเก็บอุปกรณ์</TableCell>
                                <TableCell align="center">รายละเอียดของอุปกรณ์</TableCell>
                                <TableCell align="center">พนักงานที่กรอกข้อมูล</TableCell>
                                <TableCell align="center">วันที่เพิ่ม</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {equipment.map((item:EntEquipment ) => (
                                <TableRow key={item.id}>
                                   <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.eQUIPMENTNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.equipmenttype?.eQUIPMENTTYPE}</TableCell>
                                    <TableCell align="center">{item.eQUIPMENTAMOUNT}</TableCell>
                                    <TableCell align="center">{item.edges?.classifier?.eQUIPMENTCLASSIFIER}</TableCell>
                                    <TableCell align="center">{item.edges?.zone?.eQUIPMENTZONE}</TableCell>
                                    <TableCell align="center">{item.eQUIPMENTDETAIL}</TableCell>
                                    <TableCell align="center">{item.edges?.employee?.eMPLOYEENAME}</TableCell>
                                    <TableCell align="center">{item.eQUIPMENTDATE}</TableCell>
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
                  href="./Equipment"
                >
                 ย้อนกลับ
                </Button>
            </Content>
     </Page>
    );
}