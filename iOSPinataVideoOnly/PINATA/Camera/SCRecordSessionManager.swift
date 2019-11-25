
import Foundation
import SCRecorder


let kUserDefaultsRecordSessionsKey = "RecordSessions"
class SCRecordSessionManager {
  
    private var currentSessionIdentifier: String {
      return SCRecordSessionIdentifierKey
    }
    
    private var savedSessions: [[AnyHashable : Any]] {
      get {
        return UserDefaults.standard.object(forKey: kUserDefaultsRecordSessionsKey) as? [[AnyHashable : Any]] ?? []
      }
      set(value) {
        UserDefaults.standard.setValue(value, forKey: kUserDefaultsRecordSessionsKey)
      }
    }
    
  
    public func saveRecord(recordSession: SCRecordSession?) {

      guard let currentSession = recordSession?.dictionaryRepresentation() else { return }

      var sessions = self.savedSessions
      var sessionExistingIndex = -1
      
      for i in 0...sessions.count {
        let metadata = sessions[i]
        if metadata[currentSessionIdentifier] as? String == recordSession?.identifier {
          sessionExistingIndex = i
          break
        }
      }
      
      if sessionExistingIndex == -1 {
        sessions[sessionExistingIndex] = currentSession
      } else {
        sessions.append(currentSession)
      }
      
      self.savedSessions = sessions
           
  }
    
    func removeRecord(recordSession: SCRecordSession?) {

        var sessions = self.savedSessions
      
        for var i in 0...sessions.count {
          let metadata = sessions[i]
          if metadata[currentSessionIdentifier] as? String == recordSession?.identifier {
            i -= 1
            sessions.remove(at: i)
            break
          }
        }
    
      
        self.savedSessions = sessions
    }
  
    
    func isSaved(recordSession: SCRecordSession?) -> Bool {
      
        for metadata in self.savedSessions {
          if metadata[currentSessionIdentifier] as? String == recordSession?.identifier {
            return true
          }
        }
      
        return false
    }
    
    func removeRecordSessionAtIndex(index: Int) {
      var sessions = self.savedSessions
      sessions.remove(at: index)
      self.savedSessions = sessions
        
    }
   
}
