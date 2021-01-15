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
import {
    EntEquipmenttypeEdges,
    EntEquipmenttypeEdgesFromJSON,
    EntEquipmenttypeEdgesFromJSONTyped,
    EntEquipmenttypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEquipmenttype
 */
export interface EntEquipmenttype {
    /**
     * EQUIPMENTTYPE holds the value of the "EQUIPMENTTYPE" field.
     * @type {string}
     * @memberof EntEquipmenttype
     */
    eQUIPMENTTYPE?: string;
    /**
     * 
     * @type {EntEquipmenttypeEdges}
     * @memberof EntEquipmenttype
     */
    edges?: EntEquipmenttypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntEquipmenttype
     */
    id?: number;
}

export function EntEquipmenttypeFromJSON(json: any): EntEquipmenttype {
    return EntEquipmenttypeFromJSONTyped(json, false);
}

export function EntEquipmenttypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEquipmenttype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'eQUIPMENTTYPE': !exists(json, 'EQUIPMENTTYPE') ? undefined : json['EQUIPMENTTYPE'],
        'edges': !exists(json, 'edges') ? undefined : EntEquipmenttypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntEquipmenttypeToJSON(value?: EntEquipmenttype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'EQUIPMENTTYPE': value.eQUIPMENTTYPE,
        'edges': EntEquipmenttypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}

