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
    EntPaymentEdges,
    EntPaymentEdgesFromJSON,
    EntPaymentEdgesFromJSONTyped,
    EntPaymentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPayment
 */
export interface EntPayment {
    /**
     * PAYMENTAMOUNT holds the value of the "PAYMENTAMOUNT" field.
     * @type {string}
     * @memberof EntPayment
     */
    pAYMENTAMOUNT?: string;
    /**
     * PAYMENTDATE holds the value of the "PAYMENTDATE" field.
     * @type {string}
     * @memberof EntPayment
     */
    pAYMENTDATE?: string;
    /**
     * 
     * @type {EntPaymentEdges}
     * @memberof EntPayment
     */
    edges?: EntPaymentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPayment
     */
    id?: number;
}

export function EntPaymentFromJSON(json: any): EntPayment {
    return EntPaymentFromJSONTyped(json, false);
}

export function EntPaymentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPayment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pAYMENTAMOUNT': !exists(json, 'PAYMENTAMOUNT') ? undefined : json['PAYMENTAMOUNT'],
        'pAYMENTDATE': !exists(json, 'PAYMENTDATE') ? undefined : json['PAYMENTDATE'],
        'edges': !exists(json, 'edges') ? undefined : EntPaymentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPaymentToJSON(value?: EntPayment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'PAYMENTAMOUNT': value.pAYMENTAMOUNT,
        'PAYMENTDATE': value.pAYMENTDATE,
        'edges': EntPaymentEdgesToJSON(value.edges),
        'id': value.id,
    };
}

