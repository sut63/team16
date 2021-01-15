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
    EntClassifierEdges,
    EntClassifierEdgesFromJSON,
    EntClassifierEdgesFromJSONTyped,
    EntClassifierEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClassifier
 */
export interface EntClassifier {
    /**
     * EQUIPMENTCLASSIFIER holds the value of the "EQUIPMENTCLASSIFIER" field.
     * @type {string}
     * @memberof EntClassifier
     */
    eQUIPMENTCLASSIFIER?: string;
    /**
     * 
     * @type {EntClassifierEdges}
     * @memberof EntClassifier
     */
    edges?: EntClassifierEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntClassifier
     */
    id?: number;
}

export function EntClassifierFromJSON(json: any): EntClassifier {
    return EntClassifierFromJSONTyped(json, false);
}

export function EntClassifierFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClassifier {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'eQUIPMENTCLASSIFIER': !exists(json, 'EQUIPMENTCLASSIFIER') ? undefined : json['EQUIPMENTCLASSIFIER'],
        'edges': !exists(json, 'edges') ? undefined : EntClassifierEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntClassifierToJSON(value?: EntClassifier | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'EQUIPMENTCLASSIFIER': value.eQUIPMENTCLASSIFIER,
        'edges': EntClassifierEdgesToJSON(value.edges),
        'id': value.id,
    };
}

