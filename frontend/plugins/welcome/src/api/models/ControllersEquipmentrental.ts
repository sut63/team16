/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersEquipmentrental
 */
export interface ControllersEquipmentrental {
    /**
     * 
     * @type {number}
     * @memberof ControllersEquipmentrental
     */
    employee?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersEquipmentrental
     */
    equipment?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersEquipmentrental
     */
    equipmenttype?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersEquipmentrental
     */
    member?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersEquipmentrental
     */
    rentalamount?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersEquipmentrental
     */
    rentaldate?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersEquipmentrental
     */
    returndate?: string;
}

export function ControllersEquipmentrentalFromJSON(json: any): ControllersEquipmentrental {
    return ControllersEquipmentrentalFromJSONTyped(json, false);
}

export function ControllersEquipmentrentalFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersEquipmentrental {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employee': !exists(json, 'employee') ? undefined : json['employee'],
        'equipment': !exists(json, 'equipment') ? undefined : json['equipment'],
        'equipmenttype': !exists(json, 'equipmenttype') ? undefined : json['equipmenttype'],
        'member': !exists(json, 'member') ? undefined : json['member'],
        'rentalamount': !exists(json, 'rentalamount') ? undefined : json['rentalamount'],
        'rentaldate': !exists(json, 'rentaldate') ? undefined : json['rentaldate'],
        'returndate': !exists(json, 'returndate') ? undefined : json['returndate'],
    };
}

export function ControllersEquipmentrentalToJSON(value?: ControllersEquipmentrental | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employee': value.employee,
        'equipment': value.equipment,
        'equipmenttype': value.equipmenttype,
        'member': value.member,
        'rentalamount': value.rentalamount,
        'rentaldate': value.rentaldate,
        'returndate': value.returndate,
    };
}


