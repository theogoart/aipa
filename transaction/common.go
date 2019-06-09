//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with aipa.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  trx pool interface
 * @Author: 
 * @Date:   2018-12-15
 * @Last Modified by:
 * @Last Modified time:
 */

package transaction

import (
	"github.com/aipadad/aipa/common"
	aipaErr "github.com/aipadad/aipa/common/errors"
	"github.com/aipadad/aipa/common/types"
)

type trxApplyApi interface {
	ExecuteTransaction(trx *types.Transaction)
	GetTrxErrorCode(trxHash common.Hash) aipaErr.ErrCode
}

// NewTrxApplyService is to retrieve the transaction service
func NewTrxApplyService() *TrxApplyService {
	return GetTrxApplyService()
}
