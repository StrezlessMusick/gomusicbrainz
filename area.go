/*
 * Copyright (c) 2014 Michael Wendland
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 *
 * 	Authors:
 * 		Michael Wendland <michael@michiwend.com>
 */

package gomusicbrainz

// Area represents a geographic region or settlement.
type Area struct {
	ID            string         `xml:"id,attr"`
	Type          string         `xml:"type,attr"`
	Name          string         `xml:"name"`
	SortName      string         `xml:"sort-name"`
	ISO31662Codes []ISO31662Code `xml:"iso-3166-2-code-list>iso-3166-2-code"`
	Lifespan      Lifespan       `xml:"life-span"`
	Aliases       []Alias        `xml:"alias-list>alias"`
}

// AreaSearchResponse is the response type returned by the SearchArea method.
type AreaSearchResponse struct {
	WS2ListResponse
	Areas  []*Area
	Scores ScoreMap
}

// ResultsWithScore returns a slice of Areas with a specific score.
func (r *AreaSearchResponse) ResultsWithScore(score int) []*Area {
	var res []*Area
	for k, v := range r.Scores {
		if v == score {
			res = append(res, k.(*Area))
		}
	}
	return res
}

type areaListResult struct {
	AreaList struct {
		WS2ListResponse
		Areas []struct {
			*Area
			Score int `xml:"http://musicbrainz.org/ns/ext#-2.0 score,attr"`
		} `xml:"area"`
	} `xml:"area-list"`
}
