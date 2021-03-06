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
    EntEquipment,
    EntEquipmentFromJSON,
    EntEquipmentFromJSONTyped,
    EntEquipmentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntZoneEdges
 */
export interface EntZoneEdges {
    /**
     * Equipment holds the value of the equipment edge.
     * @type {Array<EntEquipment>}
     * @memberof EntZoneEdges
     */
    equipment?: Array<EntEquipment>;
}

export function EntZoneEdgesFromJSON(json: any): EntZoneEdges {
    return EntZoneEdgesFromJSONTyped(json, false);
}

export function EntZoneEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntZoneEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'equipment': !exists(json, 'equipment') ? undefined : ((json['equipment'] as Array<any>).map(EntEquipmentFromJSON)),
    };
}

export function EntZoneEdgesToJSON(value?: EntZoneEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'equipment': value.equipment === undefined ? undefined : ((value.equipment as Array<any>).map(EntEquipmentToJSON)),
    };
}


