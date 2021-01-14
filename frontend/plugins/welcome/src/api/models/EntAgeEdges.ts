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
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntAgeEdges
 */
export interface EntAgeEdges {
    /**
     * Employee holds the value of the employee edge.
     * @type {Array<EntEmployee>}
     * @memberof EntAgeEdges
     */
    employee?: Array<EntEmployee>;
}

export function EntAgeEdgesFromJSON(json: any): EntAgeEdges {
    return EntAgeEdgesFromJSONTyped(json, false);
}

export function EntAgeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntAgeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employee': !exists(json, 'employee') ? undefined : ((json['employee'] as Array<any>).map(EntEmployeeFromJSON)),
    };
}

export function EntAgeEdgesToJSON(value?: EntAgeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employee': value.employee === undefined ? undefined : ((value.employee as Array<any>).map(EntEmployeeToJSON)),
    };
}


