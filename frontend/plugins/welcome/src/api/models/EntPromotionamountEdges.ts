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
    EntPromotion,
    EntPromotionFromJSON,
    EntPromotionFromJSONTyped,
    EntPromotionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPromotionamountEdges
 */
export interface EntPromotionamountEdges {
    /**
     * Promotion holds the value of the promotion edge.
     * @type {Array<EntPromotion>}
     * @memberof EntPromotionamountEdges
     */
    promotion?: Array<EntPromotion>;
}

export function EntPromotionamountEdgesFromJSON(json: any): EntPromotionamountEdges {
    return EntPromotionamountEdgesFromJSONTyped(json, false);
}

export function EntPromotionamountEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPromotionamountEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'promotion': !exists(json, 'promotion') ? undefined : ((json['promotion'] as Array<any>).map(EntPromotionFromJSON)),
    };
}

export function EntPromotionamountEdgesToJSON(value?: EntPromotionamountEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'promotion': value.promotion === undefined ? undefined : ((value.promotion as Array<any>).map(EntPromotionToJSON)),
    };
}


